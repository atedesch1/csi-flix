<script setup lang="ts">
import { useQuasar } from 'quasar';
import { api } from 'boot/axios';
import { Ref, ref } from 'vue';
import { matSlideshow, matMovie } from '@quasar/extras/material-icons';

const $q = useQuasar();

const name: Ref<string | null> = ref(null);
const accept = ref(null);
const movies: Ref<Array<Movie> | null> = ref(null);
interface Movie {
    Stats: {
        BoxOffice: string;
        Ratings: Array<Map<string, string>>;
        Metascore: string;
        Awards: string;
        imdbRating: string;
        imdbVotes: string;
    };
    Movie: {
        Cast: Array<string>;
        Countries: Array<string>;
        DateAdded: string;
        Directors: Array<string>;
        Description: string;
        Duration: string;
        Genres: Array<string>;
        Rating: string;
        ReleaseYear: string;
        Title: string;
        Type: string;
    };
}
async function onsubmit () {
    try{
        movies.value = (await api.get(`/movie/bytitle/${name.value}?stats=true`)).data;
        name.value = null;
        console.log(movies.value);
        $q.notify({
            color: 'green-4',
            textColor: 'white',
            icon: 'cloud_done',
            message: 'Submitted',
        });
    }catch(e) {
        console.log(e);
    }
}

function onreset () {
    name.value = null;
    accept.value = null;
}

function research() {
    movies.value = null;
}
</script>


<template>
    <div
        v-if="!movies"
        class="wrapper"
    >
        <q-form
            class="form q-gutter-lg"
            @submit="onsubmit"
            @reset="onreset"
        >
            <q-input
                v-model="name"
                filled
                label="Movie name"
                hint="Search the movie"
                lazy-rules
                :rules="[ val => val && val.length > 0 || 'Please type something']"
            />
            <div>
                <q-btn
                    label="Submit"
                    type="submit"
                    color="red-5"
                />
                <q-btn
                    label="Reset"
                    type="reset"
                    color="red-5"
                    flat
                    class="q-ml-sm"
                />
            </div>
        </q-form>
    </div>
    <div
        v-else
        class="movies q-my-xl"
    >
        <q-list
            bordered
            class="list rounded-borders q-mb-xl"
        >
            <q-expansion-item
                v-for="movie in movies"
                :key="movie.Movie.Title"
                expand-separator
                :icon="movie.Movie.Genres[0] !== 'TV Show' ? matSlideshow : matMovie"
                :label="movie.Movie.Title"
                :caption="movie.Movie.Type"
            >
                <q-card>
                    <q-card-section horizontal>
                        <q-card-section class="inner-content">
                            <q-card-section class="text-weight-bold text-h6">Movie</q-card-section>
                            <q-card-section>
                                <span class="text-weight-bold">Description: </span>
                                <span>{{ movie.Movie.Description }}</span>
                            </q-card-section>
                            <q-card-section>
                                <span class="text-weight-bold">Country: </span>
                                <span>{{ movie.Movie.Countries[0] }}</span>
                            </q-card-section>
                            <q-card-section>
                                <span class="text-weight-bold">Genre: </span>
                                <span>{{ movie.Movie.Genres[0] }}</span>
                            </q-card-section>
                            <q-card-section>
                                <span class="text-weight-bold">Rating: </span>
                                <span>{{ movie.Movie.Rating }}</span>
                            </q-card-section>
                        </q-card-section>
                        <q-card-section class="inner-content">
                            <q-card-section class="text-weight-bold text-h6">Stats</q-card-section>
                            <q-card-section v-if="movie.Stats.Awards !== ''">
                                <span class="text-weight-bold">Awards: </span>
                                <span>{{ movie.Stats.Awards }}</span>
                            </q-card-section>
                            <q-card-section v-if="movie.Stats.BoxOffice !== ''">
                                <span class="text-weight-bold">BoxOffice: </span>
                                <span>{{ movie.Stats.BoxOffice }}</span>
                            </q-card-section>
                            <q-card-section v-if="movie.Stats.Metascore !== ''">
                                <span class="text-weight-bold">Metascore: </span>
                                <span>{{ movie.Stats.Metascore }}</span>
                            </q-card-section>
                            <q-card-section v-if="movie.Stats.imdbRating !== ''">
                                <span class="text-weight-bold">imdbRating: </span>
                                <span>{{ movie.Stats.imdbRating }}</span>
                            </q-card-section>
                            <q-card-section v-if="movie.Stats.imdbVotes !== ''">
                                <span class="text-weight-bold">imdbVotes: </span>
                                <span>{{ movie.Stats.imdbVotes }}</span>
                            </q-card-section>
                        </q-card-section>
                    </q-card-section>
                </q-card>
            </q-expansion-item>
        </q-list>
        <q-btn
            label="Search another"
            color="red-5"
            @click="research"
        />
    </div>
</template>

<style scoped lang="scss">
.inner-content{
    width: 50%;
}
.list{
    width: 70%;
}
.movies {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}
.form{
    width: 30%;
}
.wrapper{
    margin-top: 10rem;
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}
</style>

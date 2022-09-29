<script setup lang="ts">
import { useQuasar } from 'quasar';
import { api } from 'boot/axios';
import { Ref, ref } from 'vue';

const $q = useQuasar();

const name: Ref<string | null> = ref(null);
const accept = ref(null);
const movies: Ref<Array<Movie> | null> = ref(null);
interface Movie {
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
}
async function onsubmit () {
    try{
        movies.value = (await api.get(`/movie/bytitle/${name.value}`)).data;
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
                label="Movie nameeee"
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
        <q-list bordered class="list rounded-borders q-mb-xl">
            <q-expansion-item
                v-for="movie in movies"
                :key="movie.Title"
                expand-separator
                icon="assessment"
                :label="movie.Title"
                :caption="movie.Countries[0]"
            >
                <q-card>
                    <q-card-section>
                        {{ movie.Description }}
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
.list{
    width: 50%;
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

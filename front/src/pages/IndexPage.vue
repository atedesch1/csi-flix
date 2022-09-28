<script setup lang="ts">
import { useQuasar } from 'quasar';
import { api } from 'boot/axios';
import { Ref, ref } from 'vue';

const $q = useQuasar();

const name: Ref<string | null> = ref(null);
const accept = ref(null);
const movies = ref(null);

async function onsubmit () {
    try{
        movies.value = (await api.get(`/movie/bytitle/${name.value}`));
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
        class="movies"
    >
        <q-btn
            label="Search another"
            color="red-5"
            @click="research"
        />
    </div>
</template>

<style scoped lang="scss">
.movies {
    width: 100vw;
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
}
.form{
    width: 30%;
}
.wrapper{
    width: 100vw;
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
}
</style>

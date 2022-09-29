<template>
    <q-layout view="hHh lpR fFf full-height">
        <q-header
            elevated
            class="bg-red-5 text-white"
        >
            <q-toolbar>
                <q-toolbar-title>
                    CSI-FLIX
                </q-toolbar-title>
            </q-toolbar>
        </q-header>
        <q-page-container class="full-width full-height">
            <router-view />
        </q-page-container>
    </q-layout>
</template>


<template>
  <div>
    <!--
      Use <q-side-link> component
      instead of <q-item> for
      internal vue-router navigation
    -->

    <q-list no-border link inset-delimiter>
      <q-list-header>Your Library</q-list-header>
      <q-side-link item to="/library/instruments" >
        <q-item-side icon="music note" ></q-item-side>
        <q-item-main label="Instruments" ></q-item-main>
      </q-side-link>
      <q-side-link item to="/library/effects" >
        <q-item-side icon="graphic eq" ></q-item-side>
        <q-item-main label="Effects" ></q-item-main>
      </q-side-link>
      <q-side-link item to="/library/samples" >
        <q-item-side icon="album" ></q-item-side>
        <q-item-main label="Samples" ></q-item-main>
      </q-side-link>
      <q-side-link item to="/library/manage" >
        <q-item-side icon="dvr" ></q-item-side>
        <q-item-main label="Manage Library" ></q-item-main>
      </q-side-link>
    </q-list>

    <q-list no-border link inset-delimiter>
      <q-list-header>Essential Links</q-list-header>
      <q-item @click="launch('https://vstguide.com')">
        <q-item-side icon="cloud" ></q-item-side>
        <q-item-main label="Online library" sublabel="vstguide.com" ></q-item-main>
      </q-item>
      <q-item @click="launch('https://vstguide.com')">
        <q-item-side icon="school" ></q-item-side>
        <q-item-main label="Docs" sublabel="docs.vstguide.com" ></q-item-main>
      </q-item>
    </q-list>
  </div>
</template>

<script>
import {
  openURL,
  QBtn,
  QIcon,
  QList,
  QListHeader,
  QItem,
  QItemSide,
  QItemMain,
  QSideLink
} from 'quasar'

export default {
  name: 'LeftSidebar',
  components: {
    QBtn,
    QIcon,
    QList,
    QListHeader,
    QItem,
    QItemSide,
    QItemMain,
    QSideLink
  },
  methods: {
    launch (url) {
      openURL(url)
    }
  }
}
</script>

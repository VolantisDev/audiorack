<template>
  <q-layout ref="layout" view="lHh Lpr fff" :left-class="{'bg-grey-2': true}">
    <q-toolbar slot="header" class="">
      <q-btn flat @click="$refs.layout.toggleLeft()">
        <q-icon name="menu" ></q-icon>
      </q-btn>

      <q-toolbar-title>
        VST Guide
        <div slot="subtitle">Manage your audio plugins, libraries, and presets, and samples.</div>
      </q-toolbar-title>
    </q-toolbar>

    <div slot="navigation">
      <router-view name="tabs"></router-view>
    </div>

    <div slot="left">
      <left-sidebar></left-sidebar>
    </div>

    <router-view name="page"></router-view>
  </q-layout>
</template>

<script>
import {
  openURL,
  QLayout,
  QToolbar,
  QToolbarTitle,
  QBtn,
  QIcon,
  QList,
  QListHeader,
  QItem,
  QItemSide,
  QItemMain,
  QTabs,
  QTab
} from 'quasar'
import LeftSidebar from './Navigation/LeftSidebar'

export default {
  name: 'Index',
  components: {
    QLayout,
    QToolbar,
    QToolbarTitle,
    QBtn,
    QIcon,
    QList,
    QListHeader,
    QItem,
    QItemSide,
    QItemMain,
    QTabs,
    QTab,
    LeftSidebar
  },
  methods: {
    launch (url) {
      openURL(url)
    }
  }
}
</script>

<style lang="stylus">

</style>

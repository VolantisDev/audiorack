import Vue from 'vue'
import Vuex from 'vuex'
import settings from './modules/settings'
import instruments from './modules/instruments'
import effects from './modules/effects'
import presets from './modules/presets'
import samples from './modules/samples'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    settings,
    instruments,
    effects,
    presets,
    samples
  }
})

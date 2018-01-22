const state = {
  settings: {}
}

const getters = {
  settings: state => state.settings
}

const mutations = {
  SET_SETTINGS (state, settings) {
    state.settings = settings
  }
}

const actions = {
  setSettings({ dispatch }, settings) {
    dispatch('SET_SETTINGS', settings)
  }
}

export default {
  state,
  getters,
  mutations,
  actions
}

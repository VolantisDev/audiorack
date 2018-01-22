const state = {
  presets: {},
  activePresetId: ''
}

const getters = {
  presets: state => state.presets,
  activePreset: state => state.presets[state.activePresetId],
}

const mutations = {
  SET_PRESETS (state, presets) {
    state.presets = presets
    state.activePresetId = ''
  },

  ADD_PRESET (state, preset) {
    state.presets[preset.id] = preset
    state.activePresetId = preset.id
  },

  EDIT_PRESET (state, preset) {
    if (state.activePresetId && state.activePresetId !== preset.id) {
      delete state.presets[state.activePresetId]
    }
    state.presets[preset.id] = preset
    state.activePresetId = preset.id
  },

  DELETE_PRESET (state) {
    if (state.activePresetId) {
      delete state.presets[state.activeEffectId]
      state.activePresetId = ''
    }
  },

  SET_ACTIVE_PRESET (state, presetId) {
    state.activePresetId = presetId
  }
}

const actions = {
  setPresets({ dispatch }, presets) {
    dispatch('SET_PRESETS', presets)
  },

  addPreset({ dispatch }, preset) {
    dispatch('ADD_PRESET', preset)
  },

  editPreset({ dispatch }, preset) {
    dispatch('EDIT_PRESET', preset)
  },

  deletePreset({ dispatch }) {
    dispatch('DELETE_PRESET')
  },

  updateActivePreset({ dispatch }, presetId) {
    dispatch('SET_ACTIVE_PRESET', presetId)
  }
}

export default {
  state,
  getters,
  mutations,
  actions
}

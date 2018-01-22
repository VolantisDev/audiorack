const state = {
  samples: {},
  activeSampleId: ''
}

const getters = {
  samples: state => state.samples,
  activeSample: state => state.samples[state.activeSampleId],
}

const mutations = {
  SET_SAMPLES (state, samples) {
    state.samples = samples
    state.activeSampleId = ''
  },

  ADD_SAMPLE (state, sample) {
    state.samples[sample.id] = sample
    state.activeSampleId = sample.id
  },

  EDIT_SAMPLE (state, sample) {
    if (state.activeSampleId && state.activeSampleId !== sample.id) {
      delete state.samples[state.activeSampleId]
    }
    state.samples[sample.id] = sample
    state.activeSampleId = sample.id
  },

  DELETE_SAMPLE (state) {
    if (state.activeSampleId) {
      delete state.samples[state.activeEffectId]
      state.activeSampleId = ''
    }
  },

  SET_ACTIVE_SAMPLE (state, sampleId) {
    state.activeSampleId = sampleId
  }
}

const actions = {
  setSamples({ dispatch }, samples) {
    dispatch('SET_SAMPLES', samples)
  },

  addSample({ dispatch }, sample) {
    dispatch('ADD_SAMPLE', sample)
  },

  editSample({ dispatch }, sample) {
    dispatch('EDIT_SAMPLE', sample)
  },

  deleteSample({ dispatch }) {
    dispatch('DELETE_SAMPLE')
  },

  updateActiveSample({ dispatch }, sampleId) {
    dispatch('SET_ACTIVE_SAMPLE', sampleId)
  }
}

export default {
  state,
  getters,
  mutations,
  actions
}

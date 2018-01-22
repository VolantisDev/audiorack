const state = {
  instruments: {},
  activeInstrumentId: ''
}

const getters = {
  instruments: state => state.instruments,
  activeInstrument: state => state.instruments[state.activeInstrumentId],
}

const mutations = {
  SET_INSTRUMENTS (state, instruments) {
    state.instruments = instruments
    state.activeInstrumentId = ''
  },

  ADD_INSTRUMENT (state, instrument) {
    state.instruments[instrument.id] = instrument
    state.activeInstrumentId = instrument.id
  },

  EDIT_INSTRUMENT (state, instrument) {
    if (state.activeInstrumentId && state.activeInstrumentId !== instrument.id) {
      delete state.instruments[state.activeInstrumentId]
    }
    state.instruments[instrument.id] = instrument
    state.activeInstrumentId = instrument.id
  },

  DELETE_INSTRUMENT (state) {
    if (state.activeInstrumentId) {
      delete state.instruments[state.activeEffectId]
      state.activeInstrumentId = ''
    }
  },

  SET_ACTIVE_INSTRUMENT (state, instrumentId) {
    state.activeInstrumentId = instrumentId
  }
}

const actions = {
  setInstruments({ dispatch }, instruments) {
    dispatch('SET_INSTRUMENTS', instruments)
  },

  addInstrument({ dispatch }, instrument) {
    dispatch('ADD_INSTRUMENT', instrument)
  },

  editInstrument({ dispatch }, instrument) {
    dispatch('EDIT_INSTRUMENT', instrument)
  },

  deleteInstrument({ dispatch }) {
    dispatch('DELETE_INSTRUMENT')
  },

  updateActiveInstrument({ dispatch }, instrumentId) {
    dispatch('SET_ACTIVE_INSTRUMENT', instrumentId)
  }
}

export default {
  state,
  getters,
  mutations,
  actions
}

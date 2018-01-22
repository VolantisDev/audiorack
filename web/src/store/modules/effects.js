const state = {
  effects: {},
  activeEffectId: ''
}

const getters = {
  effects: state => state.effects,
  activeEffect: state => state.effects[state.activeEffectId],
}

const mutations = {
  SET_EFFECTS (state, effects) {
    state.effects = effects
    state.activeEffectId = ''
  },

  ADD_EFFECT (state, effect) {
    state.effects[effect.id] = effect
    state.activeEffectId = effect.id
  },

  EDIT_EFFECT (state, effect) {
    if (state.activeEffectId && state.activeEffectId !== effect.id) {
      delete state.effects[state.activeEffectId]
    }
    state.effects[effect.id] = effect
    state.activeEffectId = effect.id
  },

  DELETE_EFFECT (state) {
    if (state.activeEffectId) {
      delete state.effects[state.activeEffectId]
      state.activeEffectId = ''
    }
  },

  SET_ACTIVE_EFFECT (state, effectId) {
    state.activeEffectId = effectId
  }
}

const actions = {
  setEffects({ dispatch }, effects) {
    dispatch('SET_EFFECTS', effects)
  },

  addEffect({ dispatch }, effect) {
    dispatch('ADD_EFFECT', effect)
  },

  editEffect({ dispatch }, effect) {
    dispatch('EDIT_EFFECT', effect)
  },

  deleteEffect({ dispatch }) {
    dispatch('DELETE_EFFECT')
  },

  updateActiveEffect({ dispatch }, effectId) {
    dispatch('SET_ACTIVE_EFFECT', effectId)
  }
}

export default {
  state,
  getters,
  mutations,
  actions
}

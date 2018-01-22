const isElectron = require('is-electron')
const firebase = require('./firebase')
const Datastore = isElectron ? require('./nedb') : firebase

// @todo Expand this to support writing to any number of datasources if configured.

export default {
  writeStateToDatastore: (state) => Datastore.writeState(state),
  readStateFromDatastore: () => Datastore.readState()
}

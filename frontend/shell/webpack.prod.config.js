const { shareAll, withModuleFederationPlugin } = require('@angular-architects/module-federation/webpack');

const NOTES_REMOTE = process.env['NOTES_REMOTE_URL'] || 'http://localhost:4201/remoteEntry.js';
const SCIM_REMOTE  = process.env['SCIM_REMOTE_URL']  || 'http://localhost:4202/remoteEntry.js';

module.exports = withModuleFederationPlugin({
  name: 'shell',
  remotes: {
    notesMFE: `notesMFE@${NOTES_REMOTE}`,
    scimMFE:  `scimMFE@${SCIM_REMOTE}`,
  },
  shared: {
    ...shareAll({ singleton: true, strictVersion: false, requiredVersion: 'auto' }),
  },
});
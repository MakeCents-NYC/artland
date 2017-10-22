//consumer key:
// PFb88Rg4l1JZ-gizJynAZ8X8RCLBT96IcRHSC7BXd8950c02!8cda42c7df3a4ecab9bd0bbf2f47f0820000000000000000
var consumer_key = "PFb88Rg4l1JZ-gizJynAZ8X8RCLBT96IcRHSC7BXd8950c02!8cda42c7df3a4ecab9bd0bbf2f47f0820000000000000000";
var key_alias = "keyalias";
var key_password = "keystorepassword";
var MasterCardAPI = require('mastercard-api-core');
var lostStolen = require('mastercard-lost-stolen');
var authentication = new MasterCardAPI.OAuth(consumer_key, "./team23-money2020-1508232975-sandbox.p12", key_alias, key_password);
MasterCardAPI.init({
    sandbox: true,
    authentication: authentication
});
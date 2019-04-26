'use strict';

var utils = require('../utils/writer.js');
var Supporterz = require('../service/SupporterzService');

module.exports.usersGET = function usersGET (req, res, next) {
  Supporterz.usersGET()
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

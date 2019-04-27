'use strict';

var utils = require('../utils/writer.js');
var Supporterz = require('../service/SupporterzService');

module.exports.v1UsersGET = function v1UsersGET (req, res, next) {
  Supporterz.v1UsersGET()
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.v1UsersIdDELETE = function v1UsersIdDELETE (req, res, next) {
  var id = req.swagger.params['id'].value;
  Supporterz.v1UsersIdDELETE(id)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.v1UsersIdGET = function v1UsersIdGET (req, res, next) {
  var id = req.swagger.params['id'].value;
  Supporterz.v1UsersIdGET(id)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.v1UsersIdPUT = function v1UsersIdPUT (req, res, next) {
  var id = req.swagger.params['id'].value;
  var body = req.swagger.params['body'].value;
  Supporterz.v1UsersIdPUT(id,body)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.v1UsersPOST = function v1UsersPOST (req, res, next) {
  var body = req.swagger.params['body'].value;
  Supporterz.v1UsersPOST(body)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

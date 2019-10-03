/**
 * ScaleShift
 * A platform for machine learning & high performance computing 
 *
 * OpenAPI spec version: 1.0.0
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.3.1
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'));
  } else {
    // Browser globals (root is window)
    if (!root.ScaleShift) {
      root.ScaleShift = {};
    }
    root.ScaleShift.AppError = factory(root.ScaleShift.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';




  /**
   * The AppError model module.
   * @module model/AppError
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>AppError</code>.
   * Application Error
   * @alias module:model/AppError
   * @class
   * @param caption {String} 
   */
  var exports = function(caption) {
    var _this = this;

    _this['caption'] = caption;




  };

  /**
   * Constructs a <code>AppError</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/AppError} obj Optional instance to populate.
   * @return {module:model/AppError} The populated <code>AppError</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('caption')) {
        obj['caption'] = ApiClient.convertToType(data['caption'], 'String');
      }
      if (data.hasOwnProperty('condition')) {
        obj['condition'] = ApiClient.convertToType(data['condition'], 'String');
      }
      if (data.hasOwnProperty('detail')) {
        obj['detail'] = ApiClient.convertToType(data['detail'], 'String');
      }
      if (data.hasOwnProperty('owner')) {
        obj['owner'] = ApiClient.convertToType(data['owner'], 'String');
      }
      if (data.hasOwnProperty('occursAt')) {
        obj['occursAt'] = ApiClient.convertToType(data['occursAt'], 'Date');
      }
    }
    return obj;
  }

  /**
   * @member {String} caption
   */
  exports.prototype['caption'] = undefined;
  /**
   * @member {String} condition
   */
  exports.prototype['condition'] = undefined;
  /**
   * @member {String} detail
   */
  exports.prototype['detail'] = undefined;
  /**
   * @member {String} owner
   */
  exports.prototype['owner'] = undefined;
  /**
   * when it happened
   * @member {Date} occursAt
   */
  exports.prototype['occursAt'] = undefined;



  return exports;
}));



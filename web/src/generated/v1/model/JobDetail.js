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
    define(['ApiClient', 'model/Job', 'model/JobFile', 'model/JobFiles', 'model/JobLog', 'model/JobLogs'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./Job'), require('./JobFile'), require('./JobFiles'), require('./JobLog'), require('./JobLogs'));
  } else {
    // Browser globals (root is window)
    if (!root.ScaleShift) {
      root.ScaleShift = {};
    }
    root.ScaleShift.JobDetail = factory(root.ScaleShift.ApiClient, root.ScaleShift.Job, root.ScaleShift.JobFile, root.ScaleShift.JobFiles, root.ScaleShift.JobLog, root.ScaleShift.JobLogs);
  }
}(this, function(ApiClient, Job, JobFile, JobFiles, JobLog, JobLogs) {
  'use strict';




  /**
   * The JobDetail model module.
   * @module model/JobDetail
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>JobDetail</code>.
   * the details of a job
   * @alias module:model/JobDetail
   * @class
   * @implements module:model/Job
   * @implements module:model/JobLogs
   * @implements module:model/JobFiles
   * @param id {String} Job ID
   * @param logs {Array.<module:model/JobLog>} 
   * @param files {Array.<module:model/JobFile>} 
   */
  var exports = function(id, logs, files) {
    var _this = this;

    Job.call(_this, id);
    JobLogs.call(_this, logs);
    JobFiles.call(_this, files);
  };

  /**
   * Constructs a <code>JobDetail</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/JobDetail} obj Optional instance to populate.
   * @return {module:model/JobDetail} The populated <code>JobDetail</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      Job.constructFromObject(data, obj);
      JobLogs.constructFromObject(data, obj);
      JobFiles.constructFromObject(data, obj);
    }
    return obj;
  }


  // Implement Job interface:
  /**
   * Job ID
   * @member {String} id
   */
exports.prototype['id'] = undefined;

  /**
   * platform
   * @member {module:model/Job.PlatformEnum} platform
   */
exports.prototype['platform'] = undefined;

  /**
   * the status of the job
   * @member {String} status
   */
exports.prototype['status'] = undefined;

  /**
   * the image ID
   * @member {String} image
   */
exports.prototype['image'] = undefined;

  /**
   * the container labels
   * @member {Array.<String>} mounts
   */
exports.prototype['mounts'] = undefined;

  /**
   * the container labels
   * @member {Array.<String>} commands
   */
exports.prototype['commands'] = undefined;

  /**
   * started unix timestamp
   * @member {Date} started
   */
exports.prototype['started'] = undefined;

  /**
   * ended unix timestamp
   * @member {Date} ended
   */
exports.prototype['ended'] = undefined;

  /**
   * A link to an external status page
   * @member {String} external_link
   */
exports.prototype['external_link'] = undefined;

  // Implement JobLogs interface:
  /**
   * @member {Array.<module:model/JobLog>} logs
   */
exports.prototype['logs'] = undefined;

  // Implement JobFiles interface:
  /**
   * @member {Array.<module:model/JobFile>} files
   */
exports.prototype['files'] = undefined;

  /**
   * @member {String} apiToken
   */
exports.prototype['apiToken'] = undefined;



  return exports;
}));


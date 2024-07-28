function _class_call_check(instance, Constructor) {
    if (!(instance instanceof Constructor)) {
        throw new TypeError("Cannot call a class as a function");
    }
}
function _define_property(obj, key, value) {
    if (key in obj) {
        Object.defineProperty(obj, key, {
            value: value,
            enumerable: true,
            configurable: true,
            writable: true
        });
    } else {
        obj[key] = value;
    }
    return obj;
}
import { Repository } from '../repositories/repository';
/**
 * Database Service class.
 *
 * @description Used to implement service classes which require a repository dependency.
 *
 * @class DBService
 * @property {Repository} repository - Repository dependency
 */ export var DBService = function DBService(connection) {
    "use strict";
    _class_call_check(this, DBService);
    _define_property(this, "repository", void 0);
    this.repository = new Repository(connection);
};

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
/**
 * Base repository class.
 *
 * @class Repository
 * @property {Connection} connection - Knex transaction client
 */ export var Repository = function Repository(connection) {
    "use strict";
    _class_call_check(this, Repository);
    _define_property(this, "connection", void 0);
    this.connection = connection;
};

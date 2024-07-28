function asyncGeneratorStep(gen, resolve, reject, _next, _throw, key, arg) {
    try {
        var info = gen[key](arg);
        var value = info.value;
    } catch (error) {
        reject(error);
        return;
    }
    if (info.done) {
        resolve(value);
    } else {
        Promise.resolve(value).then(_next, _throw);
    }
}
function _async_to_generator(fn) {
    return function() {
        var self = this, args = arguments;
        return new Promise(function(resolve, reject) {
            var gen = fn.apply(self, args);
            function _next(value) {
                asyncGeneratorStep(gen, resolve, reject, _next, _throw, "next", value);
            }
            function _throw(err) {
                asyncGeneratorStep(gen, resolve, reject, _next, _throw, "throw", err);
            }
            _next(undefined);
        });
    };
}
function _ts_generator(thisArg, body) {
    var f, y, t, g, _ = {
        label: 0,
        sent: function() {
            if (t[0] & 1) throw t[1];
            return t[1];
        },
        trys: [],
        ops: []
    };
    return g = {
        next: verb(0),
        "throw": verb(1),
        "return": verb(2)
    }, typeof Symbol === "function" && (g[Symbol.iterator] = function() {
        return this;
    }), g;
    function verb(n) {
        return function(v) {
            return step([
                n,
                v
            ]);
        };
    }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while(_)try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [
                op[0] & 2,
                t.value
            ];
            switch(op[0]){
                case 0:
                case 1:
                    t = op;
                    break;
                case 4:
                    _.label++;
                    return {
                        value: op[1],
                        done: false
                    };
                case 5:
                    _.label++;
                    y = op[1];
                    op = [
                        0
                    ];
                    continue;
                case 7:
                    op = _.ops.pop();
                    _.trys.pop();
                    continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) {
                        _ = 0;
                        continue;
                    }
                    if (op[0] === 3 && (!t || op[1] > t[0] && op[1] < t[3])) {
                        _.label = op[1];
                        break;
                    }
                    if (op[0] === 6 && _.label < t[1]) {
                        _.label = t[1];
                        t = op;
                        break;
                    }
                    if (t && _.label < t[2]) {
                        _.label = t[2];
                        _.ops.push(op);
                        break;
                    }
                    if (t[2]) _.ops.pop();
                    _.trys.pop();
                    continue;
            }
            op = body.call(thisArg, _);
        } catch (e) {
            op = [
                6,
                e
            ];
            y = 0;
        } finally{
            f = t = 0;
        }
        if (op[0] & 5) throw op[1];
        return {
            value: op[0] ? op[1] : void 0,
            done: true
        };
    }
}
import knex from 'knex';
import { APIError500 } from './error.js';
/**
 * Connection timeout milliseconds (5 seconds).
 * Note: This could be placed into ENV (I think here is fine for now).
 *
 */ var CONNECTION_TIMEOUT_MS = 5000;
/**
 * Knex client config.
 *
 */ var _knex = knex({
    client: 'pg',
    connection: {
        host: process.env.DB_HOST,
        port: Number(process.env.DB_PORT),
        database: process.env.DB_DATABASE,
        user: process.env.DB_USER_API,
        password: process.env.DB_USER_API_PASSWORD
    },
    pool: {
        min: 0,
        max: 10
    }
});
/**
 * Knex transaction client - singleton instance
 *
 */ var knexTransactionClient = _knex.transactionProvider();
/**
 * Get database connection instance, uses knex transactionProvider.
 *
 * Note: Calling this function will start the connection.
 * Connection will timeout after 5 seconds if transaction not commited / rolledback.
 *
 * @async
 * @returns {Promise<Connection>} Knex transaction provider instance
 */ export var getDBConnection = function() {
    var _ref = _async_to_generator(function() {
        var knexClient, sqlQuery;
        return _ts_generator(this, function(_state) {
            switch(_state.label){
                case 0:
                    return [
                        4,
                        knexTransactionClient()
                    ];
                case 1:
                    knexClient = _state.sent();
                    // If connection opened and not closed after 5 seconds throw error
                    setTimeout(function() {
                        if (!knexClient.isCompleted()) {
                            throw new APIError500('Transaction opened without being closed.');
                        }
                    }, CONNECTION_TIMEOUT_MS);
                    /**
   * Exectute the SQL-template-tag query with the knex transaction client.
   *
   * @async
   * @param {Sql} query - SQL statement
   * @returns {Promise<any[]>}
   */ sqlQuery = function() {
                        var _ref = _async_to_generator(function(query) {
                            var response;
                            return _ts_generator(this, function(_state) {
                                switch(_state.label){
                                    case 0:
                                        return [
                                            4,
                                            knexClient.raw(query.sql, query.values)
                                        ];
                                    case 1:
                                        response = _state.sent();
                                        return [
                                            2,
                                            response.rows
                                        ];
                                }
                            });
                        });
                        return function sqlQuery(query) {
                            return _ref.apply(this, arguments);
                        };
                    }();
                    return [
                        2,
                        {
                            knex: knexClient,
                            sql: sqlQuery,
                            commit: knexClient.commit,
                            rollback: knexClient.rollback
                        }
                    ];
            }
        });
    });
    return function getDBConnection() {
        return _ref.apply(this, arguments);
    };
}();

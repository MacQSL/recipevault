function _assert_this_initialized(self) {
    if (self === void 0) {
        throw new ReferenceError("this hasn't been initialised - super() hasn't been called");
    }
    return self;
}
function _class_call_check(instance, Constructor) {
    if (!(instance instanceof Constructor)) {
        throw new TypeError("Cannot call a class as a function");
    }
}
function _construct(Parent, args, Class) {
    if (_is_native_reflect_construct()) {
        _construct = Reflect.construct;
    } else {
        _construct = function construct(Parent, args, Class) {
            var a = [
                null
            ];
            a.push.apply(a, args);
            var Constructor = Function.bind.apply(Parent, a);
            var instance = new Constructor();
            if (Class) _set_prototype_of(instance, Class.prototype);
            return instance;
        };
    }
    return _construct.apply(null, arguments);
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
function _get_prototype_of(o) {
    _get_prototype_of = Object.setPrototypeOf ? Object.getPrototypeOf : function getPrototypeOf(o) {
        return o.__proto__ || Object.getPrototypeOf(o);
    };
    return _get_prototype_of(o);
}
function _inherits(subClass, superClass) {
    if (typeof superClass !== "function" && superClass !== null) {
        throw new TypeError("Super expression must either be null or a function");
    }
    subClass.prototype = Object.create(superClass && superClass.prototype, {
        constructor: {
            value: subClass,
            writable: true,
            configurable: true
        }
    });
    if (superClass) _set_prototype_of(subClass, superClass);
}
function _is_native_function(fn) {
    return Function.toString.call(fn).indexOf("[native code]") !== -1;
}
function _possible_constructor_return(self, call) {
    if (call && (_type_of(call) === "object" || typeof call === "function")) {
        return call;
    }
    return _assert_this_initialized(self);
}
function _set_prototype_of(o, p) {
    _set_prototype_of = Object.setPrototypeOf || function setPrototypeOf(o, p) {
        o.__proto__ = p;
        return o;
    };
    return _set_prototype_of(o, p);
}
function _type_of(obj) {
    "@swc/helpers - typeof";
    return obj && typeof Symbol !== "undefined" && obj.constructor === Symbol ? "symbol" : typeof obj;
}
function _wrap_native_super(Class) {
    var _cache = typeof Map === "function" ? new Map() : undefined;
    _wrap_native_super = function wrapNativeSuper(Class) {
        if (Class === null || !_is_native_function(Class)) return Class;
        if (typeof Class !== "function") {
            throw new TypeError("Super expression must either be null or a function");
        }
        if (typeof _cache !== "undefined") {
            if (_cache.has(Class)) return _cache.get(Class);
            _cache.set(Class, Wrapper);
        }
        function Wrapper() {
            return _construct(Class, arguments, _get_prototype_of(this).constructor);
        }
        Wrapper.prototype = Object.create(Class.prototype, {
            constructor: {
                value: Wrapper,
                enumerable: false,
                writable: true,
                configurable: true
            }
        });
        return _set_prototype_of(Wrapper, Class);
    };
    return _wrap_native_super(Class);
}
function _is_native_reflect_construct() {
    if (typeof Reflect === "undefined" || !Reflect.construct) return false;
    if (Reflect.construct.sham) return false;
    if (typeof Proxy === "function") return true;
    try {
        Boolean.prototype.valueOf.call(Reflect.construct(Boolean, [], function() {}));
        return true;
    } catch (e) {
        return false;
    }
}
function _create_super(Derived) {
    var hasNativeReflectConstruct = _is_native_reflect_construct();
    return function _createSuperInternal() {
        var Super = _get_prototype_of(Derived), result;
        if (hasNativeReflectConstruct) {
            var NewTarget = _get_prototype_of(this).constructor;
            result = Reflect.construct(Super, arguments, NewTarget);
        } else {
            result = Super.apply(this, arguments);
        }
        return _possible_constructor_return(this, result);
    };
}
var ErrorType;
(function(ErrorType) {
    ErrorType["API_ERROR"] = "ApiError";
    ErrorType["BAD_REQUEST"] = "BadRequestError";
    ErrorType["NOT_FOUND"] = "NotFoundError";
    ErrorType["SERVER_ERROR"] = "InternalServerError";
})(ErrorType || (ErrorType = {}));
/**
 * API Error class.
 *
 * @description Base class for API errors.
 *
 * @extends Error
 * @property {number} status - HTTP status code
 * @property {Array<string | object>} [errors] - Additional errors or information
 * @property {ErrorType} name - Name of error
 */ var APIError = /*#__PURE__*/ function(Error1) {
    "use strict";
    _inherits(APIError, Error1);
    var _super = _create_super(APIError);
    function APIError(message, status, errors) {
        _class_call_check(this, APIError);
        var _this;
        _this = _super.call(this, message);
        _define_property(_assert_this_initialized(_this), "status", void 0);
        _define_property(_assert_this_initialized(_this), "errors", void 0);
        _this.status = status;
        _this.errors = errors;
        _this.name = "ApiError";
        return _this;
    }
    return APIError;
}(_wrap_native_super(Error));
/**
 * Bad request 400 error.
 *
 * @description Request formatted incorrectly (body, params, query etc.)
 *
 * @extends APIError
 */ export var APIError400 = /*#__PURE__*/ function(APIError) {
    "use strict";
    _inherits(APIError400, APIError);
    var _super = _create_super(APIError400);
    function APIError400(message, errors) {
        _class_call_check(this, APIError400);
        var _this;
        _this = _super.call(this, message, 400, errors);
        _this.name = "BadRequestError";
        return _this;
    }
    return APIError400;
}(APIError);
/**
 * Not found 404 error.
 *
 * @description Requested resource not found.
 *
 * @extends APIError
 */ export var APIError404 = /*#__PURE__*/ function(APIError) {
    "use strict";
    _inherits(APIError404, APIError);
    var _super = _create_super(APIError404);
    function APIError404(message, errors) {
        _class_call_check(this, APIError404);
        var _this;
        _this = _super.call(this, message, 404, errors);
        _this.name = "NotFoundError";
        return _this;
    }
    return APIError404;
}(APIError);
/**
 * Internal server 500 error.
 *
 * @description Server failed request.
 *
 * @extends APIError
 */ export var APIError500 = /*#__PURE__*/ function(APIError) {
    "use strict";
    _inherits(APIError500, APIError);
    var _super = _create_super(APIError500);
    function APIError500(message, errors) {
        _class_call_check(this, APIError500);
        var _this;
        _this = _super.call(this, message, 500, errors);
        _this.name = "InternalServerError";
        return _this;
    }
    return APIError500;
}(APIError);

enum ErrorType {
  API_ERROR = 'ApiError',
  BAD_REQUEST = 'BadRequestError',
  NOT_FOUND = 'NotFoundError',
  SERVER_ERROR = 'InternalServerError'
}

/**
 * API Error class.
 *
 * @description Base class for API errors.
 *
 * @extends Error
 * @property {number} status - HTTP status code
 * @property {Array<string | object>} [errors] - Additional errors or information
 * @property {ErrorType} name - Name of error
 */
class APIError extends Error {
  status: number;
  errors?: Array<string | object>;

  constructor(message: string, status: number, errors?: Array<string | object>) {
    super(message);
    this.status = status;
    this.errors = errors;
    this.name = ErrorType.API_ERROR;
  }
}

/**
 * Bad request 400 error.
 *
 * @description Request formatted incorrectly (body, params, query etc.)
 *
 * @extends APIError
 */
export class APIError400 extends APIError {
  constructor(message: string, errors?: Array<string | object>) {
    super(message, 400, errors);
    this.name = ErrorType.BAD_REQUEST;
  }
}

/**
 * Not found 404 error.
 *
 * @description Requested resource not found.
 *
 * @extends APIError
 */
export class APIError404 extends APIError {
  constructor(message: string, errors?: Array<string | object>) {
    super(message, 404, errors);
    this.name = ErrorType.NOT_FOUND;
  }
}

/**
 * Internal server 500 error.
 *
 * @description Server failed request.
 *
 * @extends APIError
 */
export class APIError500 extends APIError {
  constructor(message: string, errors?: Array<string | object>) {
    super(message, 500, errors);
    this.name = ErrorType.SERVER_ERROR;
  }
}

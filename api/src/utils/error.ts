enum ErrorType {
  API_ERROR = "ApiError",
  BAD_REQUEST = "BadRequestError",
  NOT_FOUND = "NotFoundError",
  SERVER_ERROR = "InternalServerError",
}

// TODO: Update comments
class APIError extends Error {
  status: number;
  errors?: Array<string | object>;

  constructor(
    message: string,
    status: number,
    errors?: Array<string | object>,
  ) {
    super(message);
    this.status = status;
    this.errors = errors;
    this.name = ErrorType.API_ERROR;
  }
}

class APIError400 extends APIError {
  constructor(message: string, errors?: Array<string | object>) {
    super(message, 400, errors);
    this.name = ErrorType.BAD_REQUEST;
  }
}

class APIError404 extends APIError {
  constructor(message: string, errors?: Array<string | object>) {
    super(message, 404, errors);
    this.name = ErrorType.NOT_FOUND;
  }
}

class APIError500 extends APIError {
  constructor(message: string, errors?: Array<string | object>) {
    super(message, 500, errors);
    this.name = ErrorType.SERVER_ERROR;
  }
}

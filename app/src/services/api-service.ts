import { AxiosInstance } from "axios";

/**
 * @exports
 * @class APIService
 * @description Base API service - used to interface with API and backend endpoints.
 *
 */
export class APIService {
  client: AxiosInstance;

  constructor(client: AxiosInstance) {
    this.client = client;
  }
}

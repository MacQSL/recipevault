// Environment variables
export const DB_USER = process.env.DB_USER;
export const DB_USER_API = process.env.DB_USER_API;
export const DB_USER_API_PASSWORD = process.env.DB_USER_API_PASSWORD;
export const DB_SCHEMA = process.env.DB_DATABASE; // intentional - schema === database

// String sizes
export const TINY_STRING_25 = 25;
export const SMALL_STRING_100 = 100;
export const REG_STRING_250 = 250;

// Audit column comments
export const CREATE_USER_COMMENT = 'User who created the record.';
export const UPDATE_USER_COMMENT = 'User who updated the record last.';
export const CREATE_DATE_COMMENT = 'Date the record was created.';
export const UPDATE_DATE_COMMENT = 'Date the record was last updated.';

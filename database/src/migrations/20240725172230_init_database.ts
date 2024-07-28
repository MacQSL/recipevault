import type { Knex } from 'knex';
import {
  CREATE_DATE_COMMENT,
  CREATE_USER_COMMENT,
  DB_SCHEMA,
  DB_USER,
  DB_USER_API,
  DB_USER_API_PASSWORD,
  REG_STRING_250,
  SMALL_STRING_100,
  TINY_STRING_25,
  UPDATE_DATE_COMMENT,
  UPDATE_USER_COMMENT
} from '../utils/constants';

export async function up(knex: Knex): Promise<void> {
  await knex.raw(`
    -- create schema
    CREATE SCHEMA IF NOT EXISTS ${DB_SCHEMA};

    -- grant ownership to database super user
    ALTER SCHEMA ${DB_SCHEMA} OWNER TO ${DB_USER};

    -- Permantely set search path for DB_USER (postgres)
    ALTER ROLE ${DB_USER} SET search_path TO ${DB_SCHEMA}, public;

    -- create api user and grant permissions
    CREATE USER ${DB_USER_API} PASSWORD '${DB_USER_API_PASSWORD}';
    GRANT USAGE ON SCHEMA ${DB_SCHEMA} TO ${DB_USER_API};
    ALTER DEFAULT PRIVILEGES GRANT ALL ON TABLES TO ${DB_USER_API};
  `);

  //NOTE: This migration is making the assumption that the SYSTEM user is id 1.
  // The value is populated in the following migration 20240725222532_reference_data.ts

  // permission
  await knex.schema.createTable('permission', (table) => {
    table.increments('permission_id').primary().comment(`Primary identifier of the 'permission' record.`);
    table.string('name', TINY_STRING_25).notNullable().comment('Name of the permission.');
    table.string('description', REG_STRING_250).comment('Permission description or additional details.');
    // NOTE: Do we want the amount of recipes / cookbooks allowed to be stored here?
    // NOTE: Do we want to store the price of subscription here?
    table.datetime('create_date').notNullable().defaultTo('now()').comment(CREATE_DATE_COMMENT);
    table.datetime('update_date').notNullable().defaultTo('now()').comment(UPDATE_DATE_COMMENT);
    table.integer('create_user').notNullable().defaultTo(1).comment(CREATE_USER_COMMENT);
    table.integer('update_user').notNullable().defaultTo(1).comment(UPDATE_USER_COMMENT);
  });

  // user
  await knex.schema.createTable('user', (table) => {
    table.increments('user_id').primary().comment(`Primary identifier of the 'user' record.`);
    table.integer('permission_id').notNullable().comment(`Foreign key reference: 'permission.permission_id'`);
    table.string('username', TINY_STRING_25).notNullable().unique().comment('Unique username of the user.');
    table.string('description', REG_STRING_250).comment('User description or additional details.');
    table.datetime('create_date').notNullable().defaultTo('now()').comment(CREATE_DATE_COMMENT);
    table.datetime('update_date').notNullable().defaultTo('now()').comment(UPDATE_DATE_COMMENT);
    table.integer('create_user').notNullable().defaultTo(1).comment(CREATE_USER_COMMENT);
    table.integer('update_user').notNullable().defaultTo(1).comment(UPDATE_USER_COMMENT);

    table.index('permission_id').foreign('permission_id').references('permission.permission_id');
  });

  // cookbook
  await knex.schema.createTable('cookbook', (table) => {
    table.increments('cookbook_id').primary().comment(`Primary identifier of the 'cookbook' record.`);
    table.string('name', SMALL_STRING_100).notNullable().comment('Cookbook name or title.');
    table.string('description', REG_STRING_250).comment('Cookbook description.');
    table.datetime('create_date').notNullable().defaultTo('now()').comment(CREATE_DATE_COMMENT);
    table.datetime('update_date').notNullable().defaultTo('now()').comment(UPDATE_DATE_COMMENT);
    table.integer('create_user').notNullable().defaultTo(1).comment(CREATE_USER_COMMENT);
    table.integer('update_user').notNullable().defaultTo(1).comment(UPDATE_USER_COMMENT);
  });

  // user_cookbook - join table
  await knex.schema.createTable('user_cookbook', (table) => {
    table.increments('user_cookbook_id').primary().comment(`Primary identifier of the 'user_cookbook' record.`);
    table.integer('cookbook_id').notNullable().comment(`Foreign key reference: 'cookbook.cookbook_id'`);
    table.integer('user_id').notNullable().comment(`Foreign key reference: 'user.user_id'`);
    // TODO: Do we need some type of user ownership for this?
    table.datetime('create_date').notNullable().defaultTo('now()').comment(CREATE_DATE_COMMENT);
    table.datetime('update_date').notNullable().defaultTo('now()').comment(UPDATE_DATE_COMMENT);
    table.integer('create_user').notNullable().defaultTo(1).comment(CREATE_USER_COMMENT);
    table.integer('update_user').notNullable().defaultTo(1).comment(UPDATE_USER_COMMENT);

    table.unique(['cookbook_id', 'user_id']).comment('Can only assign a user to a cookbook once.');

    table.index('cookbook_id').foreign('cookbook_id').references('cookbook.cookbook_id');
    table.index('user_id').foreign('user_id').references('user.user_id');
  });

  // recipe
  await knex.schema.createTable('recipe', (table) => {
    table.increments('recipe_id').primary().comment(`Primary identifier of the 'recipe' record.`);
    table.integer('cookbook_id').notNullable().comment(`Foreign key reference: 'cookbook.cookbook_id'`);
    table.string('name', SMALL_STRING_100).notNullable().comment('Recipe name or title.');
    table.string('url').comment('External url or link of recipe.');
    table.string('description', REG_STRING_250).comment('Recipe description.');
    table.datetime('create_date').notNullable().defaultTo('now()').comment(CREATE_DATE_COMMENT);
    table.datetime('update_date').notNullable().defaultTo('now()').comment(UPDATE_DATE_COMMENT);
    table.integer('create_user').notNullable().defaultTo(1).comment(CREATE_USER_COMMENT);
    table.integer('update_user').notNullable().defaultTo(1).comment(UPDATE_USER_COMMENT);

    table.index('cookbook_id').foreign('cookbook_id').references('cookbook.cookbook_id');
  });

  // TODO: Need to add ingredient_group or alternative name ie: dry ingredients / wet ingredients etc

  // ingredient
  await knex.schema.createTable('ingredient', (table) => {
    table.increments('ingredient_id').comment(`Primary identifier of the 'ingredient' record.`);
    table.integer('recipe_id').notNullable().comment(`Foreign key reference: 'recipe.recipe_id'`);
    table.string('measurement', SMALL_STRING_100).comment('Amount of ingredient. ie: 1 tsp of sugar');
    table.string('description', REG_STRING_250).comment('Measurement additional details or description.');
    // https://medium.com/whisperarts/lexorank-what-are-they-and-how-to-use-them-for-efficient-list-sorting-a48fc4e7849f
    table.string('lexorank', TINY_STRING_25).notNullable().comment('Lexographic ordering system.');
    table.datetime('create_date').notNullable().defaultTo('now()').comment(CREATE_DATE_COMMENT);
    table.datetime('update_date').notNullable().defaultTo('now()').comment(UPDATE_DATE_COMMENT);
    table.integer('create_user').notNullable().defaultTo(1).comment(CREATE_USER_COMMENT);
    table.integer('update_user').notNullable().defaultTo(1).comment(UPDATE_USER_COMMENT);

    table.unique(['recipe_id', 'lexorank']).comment('Lexorank unique per recipe.');
    table.index('recipe_id').foreign('recipe_id').references('recipe.recipe_id');
  });
}

export async function down(_knex: Knex): Promise<void> {}

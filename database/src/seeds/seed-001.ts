import { Knex } from 'knex';

// NOTE: Static ids in the seed might be problematic

const test = 'test';

export async function seed(knex: Knex): Promise<void> {
  // Insert development users
  // Don't delete the system user populated by the migration
  await knex('user').del().whereNot({ username: 'SYSTEM' });
  await knex('user').insert([
    {
      user_id: 2,
      permission_id: knex('permission').select('permission_id').where({ name: 'Admin' }),
      username: 'MacDeluca'
    }
  ]);

  // Insert cookbooks
  await knex('cookbook').del();
  await knex('cookbook').insert([
    {
      cookbook_id: 1,
      name: 'Family Cookbook',
      description: 'Collection of family recipes'
    },
    {
      cookbook_id: 2,
      name: 'Personal Cookbook',
      description: 'Collection of personal recipes'
    }
  ]);

  // Insert user_cookbooks - join table
  await knex('user_cookbook').del();
  await knex('user_cookbook').insert([
    { cookbook_id: 1, user_id: 2 },
    { cookbook_id: 2, user_id: 2 }
  ]);

  // Insert recipes
  await knex('recipe').del();
  await knex('recipe').insert([
    {
      recipe_id: 1,
      cookbook_id: 1,
      name: 'Banana Bread',
      description: 'Moms banana bread recipe'
    },
    {
      recipe_id: 2,
      cookbook_id: 1,
      name: 'Chocolate Chip Cookies',
      description: 'Classic chocolate chip cookies with a gooey center'
    },
    {
      recipe_id: 3,
      cookbook_id: 2,
      name: 'Spaghetti Carbonara',
      description: 'Creamy pasta with pancetta and Parmesan'
    }
  ]);

  // Insert ingredients
  await knex('ingredient').del();
  await knex('ingredient').insert([
    {
      ingredient_id: 1,
      recipe_id: 1,
      measurement: '2 ripe bananas',
      description: 'Mashed',
      lexorank: 'a'
    },
    {
      ingredient_id: 2,
      recipe_id: 1,
      measurement: '1 cup sugar',
      description: 'Granulated',
      lexorank: 'b'
    },
    {
      ingredient_id: 3,
      recipe_id: 1,
      measurement: '1/2 cup butter',
      description: 'Softened',
      lexorank: 'c'
    },
    {
      ingredient_id: 4,
      recipe_id: 1,
      measurement: '1 tsp baking soda',
      description: 'Leavening agent',
      lexorank: 'd'
    },
    {
      ingredient_id: 5,
      recipe_id: 2,
      measurement: '1 cup butter',
      description: 'Softened',
      lexorank: 'a'
    },
    {
      ingredient_id: 6,
      recipe_id: 2,
      measurement: '2 cups brown sugar',
      description: 'Packed',
      lexorank: 'b'
    },
    {
      ingredient_id: 7,
      recipe_id: 2,
      measurement: '2 cups chocolate chips',
      description: 'Semi-sweet',
      lexorank: 'c'
    },
    {
      ingredient_id: 8,
      recipe_id: 3,
      measurement: '1 lb spaghetti',
      description: 'Dried pasta',
      lexorank: 'a'
    },
    {
      ingredient_id: 9,
      recipe_id: 3,
      measurement: '6 oz pancetta',
      description: 'Diced',
      lexorank: 'b'
    }
  ]);
}

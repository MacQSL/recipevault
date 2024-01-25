import type { Knex } from 'knex';

export async function up(knex: Knex): Promise<void> {
  // insert permission levels
  await knex('permission').insert([
    {
      name: 'Free',
      description: 'Upload 5 recipes for a single cookbook.'
    },
    {
      name: 'Basic',
      description: 'Unlimited recipes for a single cookbook. Can join cookbooks.'
    },
    {
      name: 'Premium',
      description: 'Unlimited recipes, cookbooks and image uploads. Can share and join cookbooks.'
    },
    {
      name: 'Admin',
      description: 'Full access with application control.'
    }
  ]);

  // insert system user account
  await knex('user').insert({
    user_id: 1, // Must be the first user ie: 1
    username: 'SYSTEM',
    permission_id: knex('permission').select('permission_id').where({ name: 'Admin' })
  });
}

export async function down(_knex: Knex): Promise<void> {}

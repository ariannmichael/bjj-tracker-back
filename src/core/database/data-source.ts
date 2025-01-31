import { DataSource } from 'typeorm';
import { User } from '../../modules/user/domain/user.entity';
import { TrainingSession } from '../../modules/training/domain/training-session.entity';
import { Technique } from '../../modules/technique/domain/technique.entity';
import { BeltProgress } from '../../modules/belt/domain/belt-progress.entity';

export const AppDataSource = new DataSource({
  type: 'postgres',
  host: process.env.DB_HOST ?? 'localhost',
  port: Number(process.env.DB_PORT) || 5432,
  username: process.env.DB_USER ?? 'postgres',
  password: process.env.DB_PASS ?? 'password',
  database: process.env.DB_NAME ?? 'bjj-tracker',
  entities: [User, TrainingSession, Technique, BeltProgress],
  migrations: ['src/migrations/*{.ts,.js}'],
  synchronize: false,
  logging: true,
});
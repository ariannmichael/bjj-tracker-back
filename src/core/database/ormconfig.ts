import { TypeOrmModule } from "@nestjs/typeorm"

export const typeOrmConfig: TypeOrmModule = {
  type: 'postgres',
  host: process.env.DB_HOST ?? 'localhost',
  port: Number(process.env.PORT) || 5432,
  username: process.env.DB_USER ?? 'postgres',
  password: process.env.DB_PASS ?? '',
  database: process.env.DB_NAME ?? 'bjj-tracker',
  entities: [],
  synchronize: true,
}
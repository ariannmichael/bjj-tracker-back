import { Module } from '@nestjs/common';
import { DatabaseModule } from './core/database/database.module';
import { UserModule } from './modules/user/user.module';
import { AuthModule } from './modules/auth/auth.module';
import { ConfigModule } from '@nestjs/config';

@Module({
  imports: [
    DatabaseModule, 
    UserModule,
    AuthModule,
    ConfigModule.forRoot({ isGlobal: true })
  ],
})
export class AppModule {}

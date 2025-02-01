import { Module } from "@nestjs/common";
import { TypeOrmModule } from "@nestjs/typeorm";
import { User } from "./domain/user.entity";
import { UserRepository } from "./infrastructure/user.repository";
import { CreateUserUseCase } from "./application/create-user.usecase";
import { UserController } from "./presentation/user.controller";

@Module({
  imports: [TypeOrmModule.forFeature([User])],
  providers: [UserRepository, CreateUserUseCase],
  controllers: [UserController],
  exports: [UserRepository]
})
export class UserModule {}
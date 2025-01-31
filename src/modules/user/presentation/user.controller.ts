import { Body, Controller, Get, Param, Post } from "@nestjs/common";
import { CreateUserUseCase } from "../application/create-user.usecase";
import { UserRepository } from "../infrastructure/user.repository";
import { User } from "../domain/user.entity";

@Controller('users')
export class UserController {
  constructor(
    private readonly createUserUseCase: CreateUserUseCase,
    private readonly userRepository: UserRepository
  ) {}

  @Get(':id')
  async getUser(@Param('id') id: string): Promise<User | null> {
    return this.userRepository.findById(id);
  }

  @Post()
  async postUser(@Body() body: { name: string, email: string, password: string }): Promise<User> {
    const { name, email, password } = body;
    return this.createUserUseCase.execute(name, email, password);
  }
}
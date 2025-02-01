import { Body, Controller, Get, Param, Patch, Post, Request, UseGuards, UsePipes, ValidationPipe } from "@nestjs/common";
import { CreateUserUseCase } from "../application/create-user.usecase";
import { UserRepository } from "../infrastructure/user.repository";
import { User } from "../domain/user.entity";
import { CreateUserDto } from "../dto/create-user.dto";
import { UpdateUserDto } from "../dto/update-user.dto";
import { JwtAuthGuard } from "src/modules/auth/jwt-auth.guard";

@Controller('users')
export class UserController {
  constructor(
    private readonly createUserUseCase: CreateUserUseCase,
    private readonly userRepository: UserRepository
  ) {}

  @Get(':id')
  @UseGuards(JwtAuthGuard)
  async getUser(@Param('id') id: string): Promise<User | null> {
    return this.userRepository.findById(id);
  }

  @Post()
  @UsePipes(new ValidationPipe({ whitelist: true, transform: true }))
  async postUser(@Body() createUserDto: CreateUserDto): Promise<User> {
    const { name, email, password } = createUserDto;
    return this.createUserUseCase.execute(name, email, password);
  }

  @Patch(':id')
  async updateUser(
    @Param('id') id: string,
    @Body() updateUserDto: UpdateUserDto
  ): Promise<User | null> {
    return this.userRepository.update(id, updateUserDto);
  }

  @Get('profile')
  @UseGuards(JwtAuthGuard)
  async getProfile(@Request() req) {
    return req.user;
  }
}
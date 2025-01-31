import { Injectable } from "@nestjs/common";
import { UserRepository } from "../infrastructure/user.repository";
import { User } from "../domain/user.entity";
import * as bcrypt from "bcrypt";

@Injectable()
export class CreateUserUseCase {
  constructor(
    private readonly userRepository: UserRepository
  ) {}

  async execute(name: string, email: string, password: string): Promise<User> {
    const existingUser = await this.userRepository.findByEmail(email);
    if (existingUser)
      throw new Error('User already exists');

    const hashedPassword = await bcrypt.hash(password, 10);

    return this.userRepository.create({ name, email, password: hashedPassword });
  }
}
import { IsEmail, IsOptional } from "class-validator";

export class UpdateUserDto {
  @IsOptional()
  name?: string;

  @IsOptional()
  @IsEmail({}, { message: 'Invalid email address' })
  email?: string;

  @IsOptional()
  @IsEmail({}, { message: 'Password must be at least 6 characters long' })
  password?: string;
}
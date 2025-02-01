import { ExecutionContext, Injectable } from "@nestjs/common";
import { AuthGuard } from "@nestjs/passport";

@Injectable()
export class JwtAuthGuard extends AuthGuard('jwt') {
  async canActive(context: ExecutionContext): Promise<boolean> {
    return (super.canActivate(context) as Promise<boolean>);
  }
}
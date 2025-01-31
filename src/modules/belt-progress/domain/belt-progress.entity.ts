import { User } from "src/modules/user/domain/user.entity";
import { Column, Entity, ManyToOne, PrimaryGeneratedColumn } from "typeorm";

@Entity()
export class BeltProgress {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  @ManyToOne(() => User, (user) => user.beltProgress, { onDelete: 'CASCADE' })
  user: User;

  @Column({ type: 'enum', enum: ['WHITE', 'BLUE', 'PURPLE', 'BROWN', 'BLACK'] })
  belt: string;

  @Column({ type: 'int', default: 0 })
  stripes: number;

  @Column({ type: 'timestamp', default: () => 'CURRENT_TIMESTAMP' })
  achievedAt: Date;
}
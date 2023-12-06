import { Module } from '@nestjs/common';
import { PrismaService } from './service/prisma.service';
import { ItemService } from './service/item.service';
import { ItemRepository } from './repository/item.repository';
import { ItemController } from './controller/item.controller';

@Module({
  imports: [],
  controllers: [ItemController],
  providers: [PrismaService, ItemService, ItemRepository],
})
export class AppModule {}

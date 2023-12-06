import { Injectable } from "@nestjs/common";
import { PrismaService } from "../service/prisma.service";

export type Item = {
    id: number;
    name: string;
    description: string;
    price?: number;
};

@Injectable()
export class ItemRepository {
    constructor(private readonly db: PrismaService) {}

    async getAllItems() {
        return await this.db.item.findMany();
    }

    async getItemById(id: number) {
        return await this.db.item.findUnique({
            where: { id },
        });
    }

    async createItem(item: Omit<Item, 'id'>) {
        return await this.db.item.create({
            data: item,
        });
    }

    async updateItem(id: number, item: Partial<Item>) {
        return await this.db.item.update({
            where: { id },
            data: item,
        });
    }

    async deleteItem(id: number) {
        return await this.db.item.delete({
            where: { id },
        });
    }

}
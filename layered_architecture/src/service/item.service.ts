import { Injectable } from "@nestjs/common";
import { Item, ItemRepository } from "src/repository/item.repository";

@Injectable()
export class ItemService {
    constructor(private readonly repository: ItemRepository) {}

    async getAllItems() {
        // Business logic here
        return await this.repository.getAllItems();
    }

    async getItemById(id: number) {
        // Business logic here
        return await this.repository.getItemById(id);
    }

    async createItem(item: Omit<Item, 'id'>) {
        // Business logic here
        return await this.repository.createItem(item);
    }

    async updateItem(id: number, item: Partial<Item>) {
        // Business logic here
        return await this.repository.updateItem(id, item);
    }

    async deleteItem(id: number) {
        // Business logic here
        return await this.repository.deleteItem(id);
    }
}


package com.example.core_net_work.model.goods;

import com.example.common.room.entitues.Product;
import com.example.core_net_work.model.BaseResult;

import java.util.ArrayList;
import java.util.List;

/**
 * @Author winiymissl
 * @Date 2024-02-23 15:26
 * @Version 1.0
 */
public class SearchResult extends BaseResult {
    SearchResultData data;

    public SearchResultData getData() {
        return data;
    }

    public void setData(SearchResultData data) {
        this.data = data;
    }

    public class SearchResultData {
        private List<Product> products = new ArrayList<>();

        public List<Product> getProducts() {
            return products;
        }

        public void setProducts(List<Product> products) {
            this.products = products;
        }
    }

}

package com.example.core_net_work.model.goods;

import com.example.core_net_work.model.BaseResult;

import java.util.List;

/**
 * @Author winiymissl
 * @Date 2024-02-28 16:26
 * @Version 1.0
 */
public class BoughtResult extends BaseResult {
    List<Temp> orders;

    public List<Temp> getOrders() {
        return orders;
    }

    public void setOrders(List<Temp> orders) {
        this.orders = orders;
    }

    public class Temp {
    }

}

package com.example.core_net_work.model.goods;

import com.example.core_net_work.model.BaseResult;

import java.util.List;

/**
 * @Author winiymissl
 * @Date 2024-02-23 15:26
 * @Version 1.0
 */
public class SearchResult extends BaseResult {
    List<ProductSimpleInfoResult.CommodityInfo> data;


    public List<ProductSimpleInfoResult.CommodityInfo> getData() {
        return data;
    }

    public void setData(List<ProductSimpleInfoResult.CommodityInfo> data) {
        this.data = data;
    }
}

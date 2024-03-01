package com.example.core_net_work.model.goods;

import com.example.common.room.UserAddress;
import com.example.core_net_work.model.BaseResult;

import java.util.List;

/**
 * @Author winiymissl
 * @Date 2024-02-29 16:49
 * @Version 1.0
 */
public class ProductInfoResult extends BaseResult {
    ProductCommodity data;

    public ProductCommodity getData() {
        return data;
    }

    public void setData(ProductCommodity data) {
        this.data = data;
    }

    public class ProductCommodity {
        ProductOneInfo commodity;

        public ProductOneInfo getCommodity() {
            return commodity;
        }

        public void setCommodity(ProductOneInfo commodity) {
            this.commodity = commodity;
        }

        public class ProductOneInfo {
            @Override
            public String toString() {
                return "ProductOneInfo{" +
                        "commodity_identity='" + commodity_identity + '\'' +
                        ", title='" + title + '\'' +
                        ", number='" + number + '\'' +
                        ", information='" + information + '\'' +
                        ", price='" + price + '\'' +
                        ", media=" + media +
                        ", address=" + address +
                        ", like_count='" + like_count + '\'' +
                        ", collect_count='" + collect_count + '\'' +
                        '}';
            }

            String commodity_identity;
            String title;
            String number;
            String information;
            String price;
            List<ProductSimpleInfoResult.CommodityInfo.MediaResult> media;
            List<UserAddress> address;
            String like_count;
            String collect_count;

            public String getCommodity_identity() {
                return commodity_identity;
            }

            public void setCommodity_identity(String commodity_identity) {
                this.commodity_identity = commodity_identity;
            }

            public String getTitle() {
                return title;
            }

            public void setTitle(String title) {
                this.title = title;
            }

            public String getNumber() {
                return number;
            }

            public void setNumber(String number) {
                this.number = number;
            }

            public String getInformation() {
                return information;
            }

            public void setInformation(String information) {
                this.information = information;
            }

            public String getPrice() {
                return price;
            }

            public void setPrice(String price) {
                this.price = price;
            }

            public List<ProductSimpleInfoResult.CommodityInfo.MediaResult> getMedia() {
                return media;
            }

            public void setMedia(List<ProductSimpleInfoResult.CommodityInfo.MediaResult> media) {
                this.media = media;
            }

//            public UserAddress getAddress() {
//                return address;
//            }
//
//            public void setAddress(UserAddress address) {
//                this.address = address;
//            }

            public String getLike_count() {
                return like_count;
            }

            public void setLike_count(String like_count) {
                this.like_count = like_count;
            }

            public String getCollect_count() {
                return collect_count;
            }

            public void setCollect_count(String collect_count) {
                this.collect_count = collect_count;
            }
        }
    }
}

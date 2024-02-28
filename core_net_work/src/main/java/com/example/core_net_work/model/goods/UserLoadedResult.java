package com.example.core_net_work.model.goods;

import com.example.core_net_work.model.BaseResult;

import java.util.List;

/**
 * @Author winiymissl
 * @Date 2024-02-26 22:08
 * @Version 1.0
 */
public class UserLoadedResult extends BaseResult {

    UserLoadedResultData data;

    public UserLoadedResultData getData() {
        return data;
    }

    public void setData(UserLoadedResultData data) {
        this.data = data;
    }

    public class UserLoadedResultData {
        UserLoadedResultDataData user_products;

        public UserLoadedResultDataData getUser_products() {
            return user_products;
        }

        public void setUser_products(UserLoadedResultDataData user_products) {
            this.user_products = user_products;
        }

        public class UserLoadedResultDataData {
//            String number;
//            String title;
//            String price;
//            String image;
            List<ProductSimpleInfoResult.CommodityInfo> Commodity;

//            public void setNumber(String number) {
//                this.number = number;
//            }

            public List<ProductSimpleInfoResult.CommodityInfo> getCommodity() {
                return Commodity;
            }

            public void setCommodity(List<ProductSimpleInfoResult.CommodityInfo> commodity) {
                this.Commodity = commodity;
            }

//            public String getTitle() {
//                return title;
//            }
//
//            public void setTitle(String title) {
//                this.title = title;
//            }
//
//            public String getPrice() {
//                return price;
//            }
//
//            public void setPrice(String price) {
//                this.price = price;
//            }
//
//            public String getImage() {
//                return image;
//            }
//
//            public void setImage(String image) {
//                this.image = image;
//            }
//
//            public String getNumber() {
//                return number;
//            }
        }
    }
}

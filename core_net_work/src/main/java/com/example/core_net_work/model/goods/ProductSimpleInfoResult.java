package com.example.core_net_work.model.goods;

import com.example.core_net_work.model.BaseResult;

import java.util.List;

/**
 * @Author winiymissl
 * @Date 2024-02-25 13:13
 * @Version 1.0
 */
public class ProductSimpleInfoResult extends BaseResult {

    public CommodityData getData() {
        return data;
    }

    public void setData(CommodityData data) {
        this.data = data;
    }

    CommodityData data;

    public class CommodityData {
        List<CommodityInfo> commodity;

        public List<CommodityInfo> getCommodity() {
            return commodity;
        }

        public void setCommodity(List<CommodityInfo> commodity) {
            this.commodity = commodity;
        }
    }


    public class CommodityInfo {

        public Long ID;
        public String commodity_identity;
        public String title;
        public String number;
        public String information;
        public float price;
        public int sold_status;
        public List<MediaResult> media;
//        public String [] media;


        public List<MediaResult> getMedia() {
            return media;
        }

        public void setMedia(List<MediaResult> media) {
            this.media = media;
        }

        public class MediaResult {
            String image;

            @Override
            public String toString() {
                return "MediaResult{" +
                        "image='" + image + '\'' +
                        '}';
            }

            public String getImage() {
                return image;
            }

            public void setImage(String image) {
                this.image = image;
            }
        }

//        public UserAddress address;

        public Long getID() {
            return ID;
        }

        @Override
        public String toString() {
            return "CommodityInfo{" + "ID=" + ID + ", commodity_identity='" + commodity_identity + '\'' + ", title='" + title + '\'' + ", number='" + number + '\'' + ", information='" + information + '\'' + ", price='" + price + '\'' + ", sold_status='" + sold_status + '\'' + ", media=" + media + '}';
        }

        public void setID(Long ID) {
            this.ID = ID;
        }

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

        public float getPrice() {
            return price;
        }

        public void setPrice(float price) {
            this.price = price;
        }

        public int getSold_status() {
            return sold_status;
        }

        public void setSold_status(int sold_status) {
            this.sold_status = sold_status;
        }

//        public MediaResult getMedia() {
//            return media;
//        }
//
//        public void setMedia(MediaResult media) {
//            this.media = media;
//        }

//        public UserAddress getAddress() {
//            return address;
//        }
//
//        public void setAddress(UserAddress address) {
//            this.address = address;
//        }
    }
}

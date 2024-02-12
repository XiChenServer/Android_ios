package com.example.mine.adpater.impl;

import com.example.mine.R;

import java.util.ArrayList;
import java.util.List;

/**
 * @Author winiymissl
 * @Date 2023-12-17 16:51
 * @Version 1.0
 */
public class MineListViewAdapterImpl {
    static List list = new ArrayList();
    int imageOne;
    String textOne;
    int more;
    static int[] image = new int[]{R.drawable.server,
            R.drawable.collect,
            R.drawable.buy,
            R.drawable.settings,};
    static String[] text = new String[]{"服务", "收藏", "购买", "设置",};

    public int getImageOne() {
        return imageOne;
    }

    public void setImageOne(int imageOne) {
        this.imageOne = imageOne;
    }

    public String getTextOne() {
        return textOne;
    }

    public void setTextOne(String textOne) {
        this.textOne = textOne;
    }

    public MineListViewAdapterImpl(int more) {

    }

    public MineListViewAdapterImpl(int imageOne, String textOne) {
        this.textOne = textOne;
        this.imageOne = imageOne;
    }

    public static List getList() {
        list.clear();
        for (int i = 0; i < text.length; i++) {
            list.add(new MineListViewAdapterImpl(image[i], text[i]));
        }
        return list;
    }
}

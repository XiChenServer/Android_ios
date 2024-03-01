package com.example.mine_pager.adapter;

import android.content.Context;
import android.view.ViewGroup;
import android.widget.ImageView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.bumptech.glide.Glide;
import com.example.mine_pager.fragment.DataBean;
import com.youth.banner.adapter.BannerAdapter;

import java.util.List;

/**
 * @Author winiymissl
 * @Date 2024-02-29 17:14
 * @Version 1.0
 */
public class MyBannerAdapter extends BannerAdapter {
    Context context;

    public MyBannerAdapter(List<DataBean> datas, Context context) {
        super(datas);
        this.context = context;
    }

    @Override
    public Object onCreateHolder(ViewGroup parent, int viewType) {
        ImageView imageView = new ImageView(parent.getContext());
        //注意，必须设置为match_parent，这个是viewpager2强制要求的
        imageView.setLayoutParams(new ViewGroup.LayoutParams(
                ViewGroup.LayoutParams.MATCH_PARENT,
                ViewGroup.LayoutParams.MATCH_PARENT));
        imageView.setScaleType(ImageView.ScaleType.CENTER_CROP);
        return new BannerViewHolder(imageView);
    }

    @Override
    public void onBindView(Object holder, Object data, int position, int size) {
        Glide.with(context).load(((DataBean) data).getUrl()).placeholder(com.example.common.R.drawable.loading).error(com.example.common.R.drawable.avatatloadfail).into(((BannerViewHolder) holder).imageView);
        // 设置边缘模糊，25 是模糊半径，3 是模糊采样

    }

    class BannerViewHolder extends RecyclerView.ViewHolder {
        public ImageView imageView;

        public BannerViewHolder(@NonNull ImageView view) {
            super(view);
            this.imageView = view;
        }
    }
}

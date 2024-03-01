package com.example.mine_pager.adapter;

import android.content.Context;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageView;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.bumptech.glide.Glide;
import com.example.common.room.entitues.ProductSimple;
import com.example.mine_pager.R;

import java.util.List;

import me.zhanghai.android.materialratingbar.MaterialRatingBar;

/**
 * @Author winiymissl
 * @Date 2024-02-12 14:01
 * @Version 1.0
 */

public class RecyclerViewAdapter extends RecyclerView.Adapter {
    //    List list;
//
//    class MyUltimateRecyclerviewHolder extends UltimateRecyclerviewViewHolder {
//        ImageView imageView;
//        TextView textView_title;
//        TextView textView_subTitle;
//        MaterialRatingBar ratingBar;
//        TextView price;
//
//        public MyUltimateRecyclerviewHolder(View itemView) {
//            super(itemView);
//            imageView = itemView.findViewById(R.id.imageView);
//            textView_title = itemView.findViewById(R.id.textView_title);
//            textView_subTitle = itemView.findViewById(R.id.textView_subTitle);
//            ratingBar = itemView.findViewById(R.id.ratingBar);
//            price = itemView.findViewById(R.id.textView_price);
//        }
//    }
//
//    public RecyclerViewAdapter(List list) {
//        this.list = list;
//    }
//
//    @Override
//    public RecyclerView.ViewHolder newFooterHolder(View view) {
//        return new MyUltimateRecyclerviewHolder(view);
//    }
//
//    @Override
//    public RecyclerView.ViewHolder newHeaderHolder(View view) {
//        return new MyUltimateRecyclerviewHolder(view);
//    }
//
//    @Override
//    public RecyclerView.ViewHolder onCreateViewHolder(ViewGroup parent) {
//        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.recycler_view_goods, parent, false);
//        return new MyUltimateRecyclerviewHolder(view);
//    }
//
//    @Override
//    public int getAdapterItemCount() {
//        return list.size();
//    }
//
//    @Override
//    public long generateHeaderId(int position) {
//        return position;
//    }
//
//    @Override
//    public void onBindViewHolder(@NonNull RecyclerView.ViewHolder holder, int position) {
//        MyUltimateRecyclerviewHolder myViewHolder = (MyUltimateRecyclerviewHolder) holder;
//        Goods goods = (Goods) list.get(position);
//        myViewHolder.ratingBar.setRating(goods.getRating());
//        myViewHolder.textView_title.setText(goods.getTitle());
//        myViewHolder.textView_subTitle.setText(goods.getSubTitle());
//        myViewHolder.imageView.setImageResource(goods.getImage());
//        myViewHolder.price.setText(goods.getPrice() + " yuan");
//    }
//
//    @Override
//    public RecyclerView.ViewHolder onCreateHeaderViewHolder(ViewGroup parent) {
//        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.recycler_view_goods, parent, false);
//        return new MyUltimateRecyclerviewHolder(view);
//    }
//
//    @Override
//    public void onBindHeaderViewHolder(RecyclerView.ViewHolder holder, int position) {
//        MyUltimateRecyclerviewHolder myViewHolder = (MyUltimateRecyclerviewHolder) holder;
//        Goods goods = (Goods) list.get(position);
//        myViewHolder.ratingBar.setRating(goods.getRating());
//        myViewHolder.textView_title.setText(goods.getTitle());
//        myViewHolder.textView_subTitle.setText(goods.getSubTitle());
//        myViewHolder.imageView.setImageResource(goods.getImage());
//        myViewHolder.price.setText(goods.getPrice() + " yuan");
//    }
    List<ProductSimple> list;
    Context context;

    public RecyclerViewAdapter(List<ProductSimple> list, Context context) {
        this.list = list;
        this.context = context;
    }

    @NonNull
    @Override
    public RecyclerView.ViewHolder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.recycler_view_goods, parent, false);
        return new MyViewHolder(view);
    }

    public class MyViewHolder extends RecyclerView.ViewHolder {
        ImageView imageView_cardView;
        TextView textView_title;
        TextView textView_subTitle;

        MaterialRatingBar ratingBar;
        TextView price;

        public MyViewHolder(@NonNull View itemView) {
            super(itemView);
            imageView_cardView = itemView.findViewById(R.id.images_cardView);
            textView_title = itemView.findViewById(R.id.textView_title);
            textView_subTitle = itemView.findViewById(R.id.textView_subTitle);
            ratingBar = itemView.findViewById(R.id.ratingBar);
            price = itemView.findViewById(R.id.textView_price);
        }
    }

    @Override
    public void onBindViewHolder(@NonNull RecyclerView.ViewHolder holder, int position) {
        MyViewHolder myViewHolder = (MyViewHolder) holder;
        ProductSimple goods = list.get(position);
        myViewHolder.ratingBar.setRating(Float.valueOf(goods.getRating()));
        myViewHolder.textView_title.setText(goods.getTitle());
        myViewHolder.textView_subTitle.setText(goods.information);
//        myViewHolder.imageView.setImageResource(R.drawable.apple);
        Glide.with(context).load(goods.getImage()).placeholder(com.example.common.R.drawable.loading).error(com.example.common.R.drawable.avatatloadfail).into(myViewHolder.imageView_cardView);
//        Glide.with(context).load(goods.getImage()).placeholder(com.example.common.R.drawable.loading).error(com.example.common.R.drawable.avatatloadfail).into(new SimpleTarget<Drawable>() {
//            @Override
//            public void onResourceReady(@NonNull Drawable resource, @Nullable Transition<? super Drawable> transition) {
//                myViewHolder.imageView_cardView.setImageDrawable(resource);
//            }
//        });
        myViewHolder.price.setText(goods.getPrice() + " yuan");
    }

    @Override
    public int getItemCount() {
        return list.size();
    }
}

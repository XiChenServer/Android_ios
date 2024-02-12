package com.example.mine_pager.adapter;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageView;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.example.mine_pager.R;
import com.example.mine_pager.adapter.model.Goods;

import java.util.List;

import me.zhanghai.android.materialratingbar.MaterialRatingBar;

/**
 * @Author winiymissl
 * @Date 2024-02-12 14:01
 * @Version 1.0
 */
public class RecyclerViewAdapter extends RecyclerView.Adapter {
    List<Goods> list;

    public RecyclerViewAdapter(List<Goods> list) {
        this.list = list;
    }

    @NonNull
    @Override
    public RecyclerView.ViewHolder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.recycler_view_goods, parent, false);
        return new MyViewHolder(view);
    }

    class MyViewHolder extends RecyclerView.ViewHolder {
        ImageView imageView;
        TextView textView_title;
        TextView textView_subTitle;

        MaterialRatingBar ratingBar;
        TextView price;

        public MyViewHolder(@NonNull View itemView) {
            super(itemView);
            imageView = itemView.findViewById(R.id.imageView);
            textView_title = itemView.findViewById(R.id.textView_title);
            textView_subTitle = itemView.findViewById(R.id.textView_subTitle);
            ratingBar = itemView.findViewById(R.id.ratingBar);
            price = itemView.findViewById(R.id.textView_price);
        }
    }

    @Override
    public void onBindViewHolder(@NonNull RecyclerView.ViewHolder holder, int position) {
        MyViewHolder myViewHolder = (MyViewHolder) holder;
        Goods goods = list.get(position);
        myViewHolder.ratingBar.setRating(goods.getRating());
        myViewHolder.textView_title.setText(goods.getTitle());
        myViewHolder.textView_subTitle.setText(goods.getSubTitle());
        myViewHolder.imageView.setImageResource(goods.getImage());
        myViewHolder.price.setText(goods.getPrice()+" yuan");
    }

    @Override
    public int getItemCount() {
        return list.size();
    }
}

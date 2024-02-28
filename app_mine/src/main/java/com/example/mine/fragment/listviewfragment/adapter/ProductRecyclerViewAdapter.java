package com.example.mine.fragment.listviewfragment.adapter;

import android.content.Context;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageView;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.bumptech.glide.Glide;
import com.example.mine.R;
import com.example.mine.fragment.listviewfragment.entity.ProductEntity;

import java.util.List;

/**
 * @Author winiymissl
 * @Date 2024-02-23 21:20
 * @Version 1.0
 */
public class ProductRecyclerViewAdapter extends RecyclerView.Adapter {

    List<ProductEntity> list;
    Context context;

    class MyViewHolder extends RecyclerView.ViewHolder {
        ImageView imageView;
        TextView textView_name;
        TextView textView_price;

        public MyViewHolder(@NonNull View itemView) {
            super(itemView);
            textView_name = itemView.findViewById(R.id.textView_product_name);
            textView_price = itemView.findViewById(R.id.textView_product_price);
            imageView = itemView.findViewById(R.id.imageView_product_image);
        }
    }

    @NonNull
    @Override
    public RecyclerView.ViewHolder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.item_recyclerview_product, parent, false);
        return new MyViewHolder(view);
    }

    public ProductRecyclerViewAdapter(List<ProductEntity> list, Context context) {
        this.list = list;
        this.context = context;
    }

    @Override
    public void onBindViewHolder(@NonNull RecyclerView.ViewHolder holder, int position) {
        MyViewHolder myViewHolder = (MyViewHolder) holder;
        ProductEntity entity = list.get(position);
        myViewHolder.textView_price.setText(entity.getPrice());
        myViewHolder.textView_name.setText(entity.getName());
        Glide.with(context).load(entity.getImage()).placeholder(com.example.common.R.drawable.loading).error(com.example.common.R.drawable.avatatloadfail).into(myViewHolder.imageView);
    }

    @Override
    public int getItemCount() {
        return list.size();
    }
}

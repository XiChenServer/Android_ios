package com.example.mine.fragment.listviewfragment.adapter;

import android.content.Context;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.bumptech.glide.Glide;
import com.example.mine.R;
import com.example.mine.fragment.listviewfragment.entity.PicRecyclerViewEntity;

import java.util.List;

/**
 * @Author winiymissl
 * @Date 2024-02-24 20:41
 * @Version 1.0
 */
public class AddPicAdapter extends RecyclerView.Adapter {

    List<PicRecyclerViewEntity> list;
    Context context;

    public AddPicAdapter(List<PicRecyclerViewEntity> list, Context context) {

        this.list = list;
        this.context = context;
    }

    @NonNull
    @Override
    public RecyclerView.ViewHolder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
        View view;
        view = LayoutInflater.from(parent.getContext()).inflate(R.layout.item_recyclerview_image, parent, false);
        return new MyViewHolderImage(view);
    }


    class MyViewHolderImage extends RecyclerView.ViewHolder {
        ImageView imageView;

        public MyViewHolderImage(@NonNull View itemView) {
            super(itemView);
            imageView = itemView.findViewById(R.id.show_pic_image);
        }
    }

    @Override
    public void onBindViewHolder(@NonNull RecyclerView.ViewHolder holder, int position) {
        if (list.size() == 0 || position == list.size()) {
            MyViewHolderImage viewHolder = (MyViewHolderImage) holder;
            Glide.with(context).load(com.example.common.R.drawable.add_count).placeholder(com.example.common.R.drawable.loading).error(com.example.common.R.drawable.avatatloadfail).into(viewHolder.imageView);
        } else {
            MyViewHolderImage viewHolder = (MyViewHolderImage) holder;
            PicRecyclerViewEntity entity = list.get(position);
            Glide.with(context).load(entity.getFile()).placeholder(com.example.common.R.drawable.loading).error(com.example.common.R.drawable.avatatloadfail).into(viewHolder.imageView);
        }
    }

    @Override
    public int getItemCount() {
        return list.size() + 1;
    }
}

package com.example.sellcowhourse.shoppingcar.adapter;

import android.content.Context;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageButton;
import android.widget.ImageView;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.bumptech.glide.Glide;
import com.example.common.room.entitues.ShoppingCarOrder;
import com.example.sellcowhourse.R;
import com.google.android.material.chip.Chip;

import java.util.List;

/**
 * @Author winiymissl
 * @Date 2024-02-22 14:18
 * @Version 1.0
 */
public class ShoppingCarOrderAdapter extends RecyclerView.Adapter {
    List<ShoppingCarOrder> list;
    Context context;

    public ShoppingCarOrderAdapter(List<ShoppingCarOrder> list, Context context) {
        this.list = list;
        this.context = context;
    }

    class MyViewHolder extends RecyclerView.ViewHolder {
        private ImageView imageView;
        private TextView name;
        private TextView price;
        private ImageButton close;
        private ImageButton add;
        private ImageButton reduce;
        private TextView count;

        public MyViewHolder(@NonNull View itemView) {
            super(itemView);
            imageView = itemView.findViewById(R.id.imageview_shopping);
            name = itemView.findViewById(R.id.textView_name);
            price = itemView.findViewById(R.id.textView_price);
            close = itemView.findViewById(R.id.imageButton_close);
            add = itemView.findViewById(R.id.chip_add);
            reduce = itemView.findViewById(R.id.chip_reduce);
            count = itemView.findViewById(R.id.textView_count);
        }
    }

    @NonNull
    @Override
    public RecyclerView.ViewHolder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.item_shopping_car_ordor, parent, false);
        return new MyViewHolder(view);
    }

    @Override
    public void onBindViewHolder(@NonNull RecyclerView.ViewHolder holder, int position) {
        ShoppingCarOrder order = list.get(position);
        MyViewHolder viewHolder = (MyViewHolder) holder;
        Glide.with(context).load(order.getImage()).error(com.example.common.R.drawable.avatatloadfail).into(viewHolder.imageView);
        viewHolder.price.setText("￥ " + String.valueOf(order.getPrice()));
        viewHolder.name.setText(String.valueOf(order.getName()));
        viewHolder.count.setText(String.valueOf(order.getCount()));
        viewHolder.add.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                int count = order.getCount();
                order.setCount(++count);
                notifyItemChanged(position);
            }
        });
        viewHolder.reduce.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                int count = order.getCount();
                if (count > 0) {
                    order.setCount(--count);
                    notifyItemChanged(position);
                }
            }
        });
        viewHolder.close.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                list.remove(position);
                notifyDataSetChanged();
            }
        });
    }

    @Override
    public int getItemCount() {
        return list.size();
    }
}

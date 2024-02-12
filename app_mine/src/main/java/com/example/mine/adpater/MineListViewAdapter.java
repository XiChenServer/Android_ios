package com.example.mine.adpater;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.BaseAdapter;
import android.widget.ImageView;
import android.widget.TextView;

import com.example.mine.adpater.impl.MineListViewAdapterImpl;
import com.example.mine.R;

import java.util.List;

/**
 * @Author winiymissl
 * @Date 2023-12-17 16:50
 * @Version 1.0
 */
public class MineListViewAdapter extends BaseAdapter {
    public MineListViewAdapter(List list) {
        this.list = list;
    }

    List<MineListViewAdapterImpl> list;

    @Override
    public int getCount() {
        return list.size();
    }

    @Override
    public Object getItem(int position) {
        return list.get(position);
    }

    @Override
    public long getItemId(int position) {
        return position;
    }

    @Override
    public View getView(int position, View convertView, ViewGroup parent) {
        View view = null;
        if (convertView == null) {
            view = LayoutInflater.from(parent.getContext()).inflate(R.layout.listview_item_mine, parent, false);
        } else {
            view = convertView;
        }
        list = MineListViewAdapterImpl.getList();

        ImageView imageView_1 = view.findViewById(R.id.imageView_listview_1);
        ImageView imageView_2 = view.findViewById(R.id.imageView_listview_2);
        TextView textView = view.findViewById(R.id.textView_listview);
        imageView_1.setImageResource(list.get(position).getImageOne());
        imageView_2.setImageResource(R.drawable.right);
        textView.setText(list.get(position).getTextOne());

        return view;
    }
}

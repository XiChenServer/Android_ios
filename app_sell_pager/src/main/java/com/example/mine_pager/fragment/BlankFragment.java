package com.example.mine_pager.fragment;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.MenuItem;
import android.view.View;
import android.view.ViewGroup;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.recyclerview.widget.GridLayoutManager;

import com.example.mine_pager.R;
import com.example.mine_pager.adapter.RecyclerViewAdapter;
import com.example.mine_pager.adapter.model.Goods;
import com.example.mine_pager.databinding.FragmentBlankBinding;
import com.kennyc.bottomsheet.BottomSheetListener;
import com.kennyc.bottomsheet.BottomSheetMenuDialogFragment;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;


public class BlankFragment extends Fragment {
    FragmentBlankBinding binding;

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_blank, container, false);
        binding = FragmentBlankBinding.bind(view);
        binding.sellPagerRecyclerView.setLayoutManager(new GridLayoutManager(getActivity(), 2));
        List<Goods> list = new ArrayList<>();
        list.add(new Goods(R.drawable.apple, "苹果", "大苹果", 4.5f, 20));
        list.add(new Goods(R.drawable.banana, "香蕉", "一个香蕉", 3.2f, 30));
        list.add(new Goods(R.drawable.strawberry, "草莓", "大草莓", 5.0f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 40));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 40));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
        RecyclerViewAdapter recyclerViewAdapter = new RecyclerViewAdapter(list);
        binding.sellPagerRecyclerView.setAdapter(recyclerViewAdapter);
        binding.imageButtonSort.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                new BottomSheetMenuDialogFragment.Builder(getActivity()).setSheet(R.menu.menu_sort).setTitle("Sort").setListener(new BottomSheetListener() {
                    @Override
                    public void onSheetShown(@NonNull BottomSheetMenuDialogFragment bottomSheetMenuDialogFragment, @Nullable Object o) {

                    }

                    @Override
                    public void onSheetItemSelected(@NonNull BottomSheetMenuDialogFragment bottomSheetMenuDialogFragment, @NonNull MenuItem menuItem, @Nullable Object o) {
                        if (menuItem.getItemId() == R.id.sort_max_min) {
                            //从高到低
                            list.sort(new Comparator<Goods>() {
                                @Override
                                public int compare(Goods o1, Goods o2) {
                                    if (o1.getPrice() - o2.getPrice() < 0) {
                                        return 1;
                                    }
                                    return -1;
                                }
                            });
                            recyclerViewAdapter.notifyDataSetChanged();
                        } else if (menuItem.getItemId() == R.id.sort_min_max) {
                            //从低到高
                            list.sort(new Comparator<Goods>() {
                                @Override
                                public int compare(Goods o1, Goods o2) {
                                    if (o1.getPrice() - o2.getPrice() > 0) {
                                        return 1;
                                    }
                                    return -1;
                                }
                            });
                            recyclerViewAdapter.notifyDataSetChanged();
                        }
                    }

                    @Override
                    public void onSheetDismissed(@NonNull BottomSheetMenuDialogFragment bottomSheetMenuDialogFragment, @Nullable Object o, int i) {

                    }
                }).show(getActivity().getSupportFragmentManager());
            }
        });
        binding.imageButtonFilter.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                //直接新建一个碎片
            }
        });
        binding.sellPagerToolbar.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                //直接新建一个搜索碎片

            }
        });
        return view;
    }
}
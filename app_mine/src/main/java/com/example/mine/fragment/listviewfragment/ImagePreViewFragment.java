package com.example.mine.fragment.listviewfragment;

import android.os.Bundle;
import android.view.GestureDetector;
import android.view.LayoutInflater;
import android.view.MotionEvent;
import android.view.View;
import android.view.ViewGroup;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.fragment.app.FragmentManager;
import androidx.fragment.app.FragmentTransaction;

import com.bumptech.glide.Glide;
import com.example.mine.R;
import com.example.mine.databinding.FragmentPreviewImageBinding;
import com.example.mine.fragment.listviewfragment.eventbus.SendMessageEvent;

import org.greenrobot.eventbus.EventBus;

/**
 * @Author winiymissl
 * @Date 2024-02-25 19:48
 * @Version 1.0
 */
public class ImagePreViewFragment extends Fragment {
    FragmentPreviewImageBinding binding;

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View view = LayoutInflater.from(getActivity()).inflate(R.layout.fragment_preview_image, container, false);
        binding = FragmentPreviewImageBinding.bind(view);
        binding.transparentViewCollect.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {

            }
        });
        Bundle bundle = getArguments();
        String path = bundle.getString("path");
        int position = bundle.getInt("position");
        Glide.with(getActivity()).load(path).error(com.example.common.R.drawable.avatatloadfail).into(binding.imageviewPreviewImage);
        GestureDetector gestureDetector = new GestureDetector(getActivity(), new MyGestureListener());
        binding.chip.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                EventBus.getDefault().post(new SendMessageEvent(position, true));
                FragmentManager fragmentManager = getActivity().getSupportFragmentManager();
                FragmentTransaction fragmentTransaction = fragmentManager.beginTransaction();
                fragmentTransaction.setCustomAnimations(com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right, com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right);
                fragmentManager.popBackStack();
                fragmentTransaction.commit();
            }
        });
        binding.imageviewPreviewImage.setOnTouchListener(new View.OnTouchListener() {
            @Override
            public boolean onTouch(View v, MotionEvent event) {
                gestureDetector.onTouchEvent(event);
                return true;
            }
        });
        return binding.getRoot();
    }

    private class MyGestureListener extends GestureDetector.SimpleOnGestureListener {

        private static final float MIN_SCALE_FACTOR = 1.0f;
        private static final float MAX_SCALE_FACTOR = 2.0f;

        private float scaleFactor = 1.0f;
        private float lastX;
        private float lastY;

        @Override
        public boolean onDoubleTap(@NonNull MotionEvent e) {
            float focusX = e.getX();
            float focusY = e.getY();
            scaleFactor = scaleFactor > 1.0f ? MIN_SCALE_FACTOR : MAX_SCALE_FACTOR;
            // 计算相对于缩放中心点的偏移量
            float offsetX = scaleFactor;
            float offsetY = scaleFactor;

            binding.imageviewPreviewImage.setPivotX(focusX);
            binding.imageviewPreviewImage.setPivotY(focusY);
            binding.imageviewPreviewImage.setScaleX(offsetX);
            binding.imageviewPreviewImage.setScaleY(offsetY);

            return true;
        }

//        @Override
//        public boolean onDown(MotionEvent e) {
//            lastX = e.getX();
//            lastY = e.getY();
//            return true;
//        }

//        @Override
//        public boolean onScroll(@NonNull MotionEvent e1, @NonNull MotionEvent e2, float distanceX, float distanceY) {
//            float deltaX = e2.getX() - lastX;
//            float deltaY = e2.getY() - lastY;
//
//
//            float newX = binding.imageviewPreviewImage.getX() + deltaX;
//            float newY = binding.imageviewPreviewImage.getY() + deltaY;
//
//            binding.imageviewPreviewImage.setX(newX);
//            binding.imageviewPreviewImage.setY(newY);
//
//            lastX = e2.getX();
//            lastY = e2.getY();
//
//            return true;
//        }
    }
}

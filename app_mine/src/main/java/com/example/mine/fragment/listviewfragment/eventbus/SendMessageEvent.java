package com.example.mine.fragment.listviewfragment.eventbus;

/**
 * @Author winiymissl
 * @Date 2024-02-25 20:46
 * @Version 1.0
 */
public class SendMessageEvent {
    private int position;
    private boolean isDelete;

    public int getPosition() {
        return position;
    }

    public void setPosition(int position) {
        this.position = position;
    }

    public boolean isDelete() {
        return isDelete;
    }

    public SendMessageEvent(int position, boolean isDelete) {
        this.position = position;
        this.isDelete = isDelete;
    }

    public void setDelete(boolean delete) {
        isDelete = delete;
    }
}


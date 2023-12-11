//
//  SettingTableViewCell.m
//  three
//
//  Created by Huc on 2023/12/10.
//

#import "SettingTableViewCell.h"

@implementation SettingTableViewCell

- (instancetype)initWithStyle:(UITableViewCellStyle)style reuseIdentifier:(NSString *)reuseIdentifier {
    self = [super initWithStyle:style reuseIdentifier:reuseIdentifier];
    
    self.imageBtn = [UIButton buttonWithType:UIButtonTypeCustom];
    self.mainLabel = [[UILabel alloc] init];
    
    [self.contentView addSubview:self.imageBtn];
    [self.contentView addSubview:self.mainLabel];
    
    return self;
}

- (void)layoutSubviews {
    self.imageBtn.frame = CGRectMake(15, 15, 30, 30);
    self.mainLabel.frame = CGRectMake(60, 15, 200, 30);
    
    self.mainLabel.textColor = UIColor.blackColor;
    
    self.imageBtn.backgroundColor = UIColor.grayColor;
//    self.mainLabel.backgroundColor = UIColor.grayColor;
    
    self.backgroundColor = UIColor.whiteColor;
}

- (void)awakeFromNib {
    [super awakeFromNib];
    // Initialization code
}

- (void)setSelected:(BOOL)selected animated:(BOOL)animated {
    [super setSelected:selected animated:animated];

    // Configure the view for the selected state
}

@end

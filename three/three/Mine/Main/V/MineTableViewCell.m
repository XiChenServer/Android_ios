//
//  MineTableViewCell.m
//  three
//
//  Created by Huc on 2023/12/10.
//

#import "MineTableViewCell.h"

@implementation MineTableViewCell

- (instancetype)initWithStyle:(UITableViewCellStyle)style reuseIdentifier:(NSString *)reuseIdentifier {
    self = [super initWithStyle:style reuseIdentifier:reuseIdentifier];
    
    self.imageBtn = [UIButton buttonWithType:UIButtonTypeCustom];
    self.mainLabel = [[UILabel alloc] init];
    self.subLabel = [[UILabel alloc] init];
    
    [self.contentView addSubview:self.imageBtn];
    [self.contentView addSubview:self.mainLabel];
    [self.contentView addSubview:self.subLabel];
    
    return self;
}

- (void)layoutSubviews {
    self.imageBtn.frame = CGRectMake(20, 20, 70, 70);
    self.mainLabel.frame = CGRectMake(110, 20, 200, 40);
    self.subLabel.frame = CGRectMake(110, 60, 200, 30);
    
    self.mainLabel.text = @"Lai'";
    self.subLabel.text = @"tel:13228059782";
    
    self.mainLabel.font = [UIFont fontWithName:@"Arial-BoldMT" size:20];
    self.subLabel.font = [UIFont systemFontOfSize:16];
    self.mainLabel.textColor = UIColor.blackColor;
    self.subLabel.textColor = UIColor.grayColor;
    
    self.imageBtn.backgroundColor = UIColor.grayColor;
}

- (void)awakeFromNib {
    [super awakeFromNib];
    // Initialization code
}

- (void)setSelected:(BOOL)selected animated:(BOOL)animated {
    // Configure the view for the selected state
}

@end
